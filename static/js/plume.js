const $ = document.querySelector.bind(document);
const $$ = document.querySelectorAll.bind(document);

const enterForm = $('#enterForm');
const messageForm = $('#messageForm');

const TYPE = {
    Hello: 0,
    Text: 1,
}

function show(el) {
    el.style.display = 'flex';
}

function hide(el) {
    el.style.display = 'none';
}

function scrollToEnd() {
    const container = $('.messageList');
    container.scrollTop = container.getBoundingClientRect().height;
}

function logMessage(user, text) {
    const item = document.createElement('li');
    item.classList.add('messageItem');

    const userSpan = document.createElement('span');
    userSpan.classList.add('user');
    userSpan.textContent = `@${user}:`;
    userSpan.style.color = colorizeString(user);
    item.appendChild(userSpan);

    item.appendChild(document.createTextNode(' '));

    const textSpan = document.createElement('span');
    textSpan.classList.add('text');
    textSpan.textContent = text;
    item.appendChild(textSpan);

    $('.messageList').appendChild(item);
    scrollToEnd();
}

let conn = null;

function connect(name, email, cb) {
    conn = new WebSocket(`ws://${window.location.host}/connect`);
    conn.addEventListener('open', evt => {
        conn.send(JSON.stringify({
            type: TYPE.Hello,
            text: `${name}\n${email}`,
        }))

        if (cb) cb();
    });
    conn.addEventListener('message', evt => {
        const message = JSON.parse(evt.data);
        logMessage(message.user.name, message.text);
    });
    conn.addEventListener('error', evt => {
        console.log('WebSocket error:', evt);
    });
}

function send(text) {
    if (conn === null) {
        return;
    }

    conn.send(JSON.stringify({
        type: TYPE.Text,
        text: text,
    }));
}

function close() {
    if (conn === null) {
        return;
    }

    conn.close();
    conn = null;
}

function colorizeString(s) {
    let hash = 0;
    for (let i = 0, len = s.length; i < len; i ++) {
        let ch = s.charCodeAt(i);
        hash = ((hash << 5) - hash) + ch;
        hash = hash & hash;
    }
    return `hsl(${Math.abs(hash % 360)}, 90%, 36%)`
}

enterForm.addEventListener('submit', evt => {
    evt.preventDefault();

    const name = enterForm.querySelector('[name="name"]').value;
    const email = enterForm.querySelector('[name="email"]').value;

    if (!name || !email) {
        return
    }

    connect(name, email, () => {
        hide(enterForm);
        show(messageForm);
        messageForm.querySelector('[name="text"]').focus();
    });
});

messageForm.addEventListener('submit', evt => {
    evt.preventDefault();

    const textInput = messageForm.querySelector('[name="text"]');
    const text = textInput.value;

    if (!text.trim()) {
        return;
    }

    send(text);
    textInput.value = '';
});

hide(messageForm);
