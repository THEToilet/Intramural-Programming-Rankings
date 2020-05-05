FROM golang

WORKDIR /home/rank/


CMD ["cd", "/home/rank/"]

CMD ["ls", "/home/rank"]
#バイナリ実行
CMD ["/home/uploader/Intramural-Programming-Rankings"]