FROM postgres:14-alpine

# configura usuario, senha e db
ENV POSTGRES_USER=tcc
ENV POSTGRES_PASSWORD=tcc
ENV POSTGRES_DB=tcc

# copia o script de inicializacao
COPY create.sql /docker-entrypoint-initdb.d/

# expoe a porta 5432
EXPOSE 5432
