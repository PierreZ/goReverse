# goReverse
A small reverse-proxy implementation written in Golang

## Why
This small code is trying to simplify the headache I have each time I'm trying to configure Apache2 or nginx for a simple task (Reverse-proxy). I found stupid to just copy a config file that you're not capable of writing it by yourself. And because I like go code.

## Usage

```
goReverse --p=8080 --d="http://your_dest:port"
```