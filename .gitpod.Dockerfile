FROM gitpod/workspace-postgres

# Install the protobuf compiler.
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip -O proto.zip \
  && sudo unzip -j proto.zip bin/protoc -d /usr/local/bin \
  && rm proto.zip
