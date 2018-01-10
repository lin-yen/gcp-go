FROM golang

# 因為要使用 go install，所以要放到指定目錄 /go/src/ 底下
ADD demo/ /go/src/test

# go install 路徑規則
# test -> /go/src/test
# ./test -> ~/test
RUN go install test

# go install 編譯好後會放置於 /go/bin 下
ENTRYPOINT /go/bin/test

# 開啟 4000 port 給機器內部使用
EXPOSE 4000
