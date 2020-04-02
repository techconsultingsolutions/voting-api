FROM golang as runtime

WORKDIR /app

COPY . ./

ENTRYPOINT ["/bin/bash"]