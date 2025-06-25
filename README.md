# auto-vpn
Simple program to auto connect to .ovpn files using `openvpn`

<br>
[Usage](#Usage) <span>&nbsp;•&nbsp;</span> [Install](#Install)

## Usage
- Run `auto-vpn` with `-p` followed by a file path to listen on
```bash
auto-vpn -p ~/Downloads
```
- Use `-clean` to remove old .ovpn files before listening
```bash
auto-vpn -p ~/Downloads -clean
```

## Install
Download pre-built binary for your system here [Releases](https://github.com/CoreyRobinsonDev/auto-vpn/releases).

### Compiling from Source
- Clone this repository
```bash
git clone https://github.com/CoreyRobinsonDev/auto-vpn.git
```
- Create **auto-vpn** binary
```bash
cd auto-vpn
go build -o /usr/local/bin
```
- Confirm that the program was built successfully
```bash
auto-vpn -v
```
## Running without Password
This program is meant to run without intervention in the background. However `openvpn` is ran with `sudo` and `sudo` requires you to enter your user password. <br>
To get around this do the following:

- Edit `/etc/sudoers` with `visudo`
```bash
sudo visudo
```
- Add the following with your username and path to `openvpn`
```bash
[your_username] ALL=(ALL) NOPASSWD: [/path/to/openvpn]
```
## License
[The Unlicense](./LICENSE)
