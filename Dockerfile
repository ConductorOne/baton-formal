FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-formal"]
COPY baton-formal /