FROM ubuntu:latest


# instala dependencias
RUN apt-get update \ 
    && DEBIAN_FRONTEND=noninteractive apt-get -y install python3 python3-pip cron \
    # Remove package lists for smaller image sizes
    && rm -rf /var/lib/apt/lists/* \
    && which cron \
    && rm -rf /etc/cron.*/*

# copia o crontab
COPY microhorario_cron.txt /etc/cron.d/microhorario_cron
RUN chmod 0644 /etc/cron.d/microhorario_cron

# instala o cron
RUN crontab /etc/cron.d/microhorario_cron

# copy o entrypath (para importar envs)
COPY entrypoint.sh /entrypoint.sh

# copia o codigo
COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt
COPY *.py .
RUN chmod 0744 *.py

# executa o cron
CMD ["cron", "-f"]