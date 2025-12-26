export const useExemplo = () => {
    const contador = useState('contador', () => 0)

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