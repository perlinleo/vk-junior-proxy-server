# HTTP-proxy

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
* Домашнее задание должно представлять из себя Docker контейнер или Docker compose файл
* Программа должна слушать на 8080 порту и проксировать HTTP трафик
* Несмотря на то, что задание называется HTTP-прокси, стоит работать с TCP сокетами

## Docker

```shell
docker build -t proxy .
docker run -d -p 8080:8080 -t proxy
```
## Usage

```shell
curl -x http://127.0.0.1:8080 http://vk.com
```
