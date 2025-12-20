// Isso equivale ao seu Service em Go
export const useExemplo = () => {
    // Estado (State)
    const contador = useState('contador', () => 0)

    // Método (Regra de negócio)
    const incrementar = () => {
        if (contador.value < 10) {
            contador.value++
        } else {
            alert('Regra de negócio: Máximo é 10!')
        }
    }

    return {
        contador,
        incrementar
    }
}