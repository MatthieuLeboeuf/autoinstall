# :eagle: autoinstall
App permettant de simplifier l'installation et la configuration d'un serveur linux

[![Build Status](https://img.shields.io/github/workflow/status/MatthieuLeboeuf/autoinstall/Build)](https://github.com/MatthieuLeboeuf/autoinstall/actions)
[![Open Issues](https://img.shields.io/github/issues-raw/MatthieuLeboeuf/autoinstall)](https://github.com/MatthieuLeboeuf/autoinstall/issues)
[![License](https://img.shields.io/github/license/MatthieuLeboeuf/autoinstall)](https://github.com/MatthieuLeboeuf/autoinstall/blob/main/LICENSE)

## Fonctionnalités

- Installation de plusieurs paquets (nodejs, yarn, php, etc)
- Installation de LAMP et LEMP
- Installation de certificats ssl (Let's Encrypt, Cloudflare)

## Installation
```
curl -s https://deb.matthieul.dev/gpg.key | gpg --dearmor | tee /usr/share/keyrings/matthieul.gpg >/dev/null
echo 'deb [signed-by=/usr/share/keyrings/matthieul.gpg] https://deb.matthieul.dev stable main' > /etc/apt/sources.list.d/autoinstall.list
apt update
apt install autoinstall -y
```