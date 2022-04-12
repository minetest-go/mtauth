FROM scratch
COPY mtauth /bin/mtauth
EXPOSE 8080
ENTRYPOINT ["/bin/mtauth"]
