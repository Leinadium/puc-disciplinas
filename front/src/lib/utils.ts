/**
 * Função simples de pluralização
 * @param n numero de elementos
 * @param text texto a ser pluralizado
 * @returns 
 */
export function pluralize(n: number, text: string): string {
    if (n == 1) return text;
    if (text.endsWith("s")) return text + "es";
    if (text.endsWith("ão")) return text.substring(0, text.length - 2) + "ões";
    return text + "s";
} 