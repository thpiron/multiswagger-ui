FROM scratch

COPY dist/swaggerui /swaggerui

CMD ["/swaggerui"]