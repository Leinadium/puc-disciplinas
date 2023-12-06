from random import sample, randint
from locust import HttpUser, task, between

escolhas: list = [
    {"disciplina": "INF1950", "turma": "3WF"},
    {"disciplina": "ENG1133", "turma": "3VB"},
    {"disciplina": "MQI2957", "turma": "32A"},
    {"disciplina": "ENG4402", "turma": "3VA"},
    {"disciplina": "ENG1443", "turma": "3VA"},
    {"disciplina": "ENG1438", "turma": "3VA"}
]


def recomendacao_aleatoria():
    return sample(escolhas, randint(0, len(escolhas) - 1))


class Usuario(HttpUser):
    wait_time = between(1, 5)

    def on_start(self):
        # se autentica
        usuario = '0000000'
        senha = '123123'
        self.client.post(f"/login?usuario={usuario}&senha={senha}")

        return super().on_start()

    @task
    def recomendacao(self):
        payload = {"escolhas": recomendacao_aleatoria()}
        self.client.post("/recomendacao", json=payload)
