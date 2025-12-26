<template>
  <div class="p-10 flex flex-col gap-4 items-start">
    <h1 class="text-3xl font-bold">Painel do CRM</h1>
    
    <p class="text-muted-foreground">
      Contador atual: {{ contador }}
    </p>

    <Button @click="incrementar">
      Aumentar Score
    </Button>
  </div>
</template>

<script setup lang="ts">
const { contador, incrementar } = useExemplo()
</script>