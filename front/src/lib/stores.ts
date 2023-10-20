import { writable } from "svelte/store";
import type { Writable } from "svelte/store";
import type { DisciplinaBasica } from "../types/disciplinas";

/** o nome do usuario, se estiver logado, ou nulo */
export const userStore: Writable<string | null> = writable(null);

/** mapa de codigo -> conteudo da disciplina. Contem as informacoes necessarias de todas as disciplinas */
export const disciplinasStore: Writable<Map<string, DisciplinaBasica>> = writable(new Map());
/** conjunto de codigo de disciplinas que o usuario ainda nao cursou */
export const faltaCursarStore: Writable<Set<string>> = writable(new Set());
/** conjunto de codigo de disciplinas que o usuario pode cursar */
export const podeCursarStore: Writable<Set<string>> = writable(new Set());