FROM scratch

COPY --chown=65534:0 ttk-exporter /
VOLUME /tmp

USER 65534

ENTRYPOINT ["/ttk-exporter"]