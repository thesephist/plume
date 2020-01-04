# Plume

Tiny in-memory chat server, running at [plume.chat](https://plume.chat).

## Deploy

Deployment is managed by systemd. Copy the `plume.service` file to `/etc/systemd/system/plume.service` and update:

- replace `plume-user` with your Linux user
- replace `/home/plume-user/plume` with your working directory

Then start Plume as a service:

```sh
# systemctl start plume
```
