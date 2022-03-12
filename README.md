# autoinstall
App permettant de simplifier l'installation et la configuration d'un serveur linux

Pour installer autoinstall sur votre machine :
```
curl -s https://deb.matthieul.dev/gpg.key | gpg --dearmor | tee /usr/share/keyrings/matthieul.gpg >/dev/null
echo 'deb [signed-by=/usr/share/keyrings/matthieul.gpg] https://deb.matthieul.dev stable main' > /etc/apt/sources.list.d/autoinstall.list
apt update
apt install autoinstall
```