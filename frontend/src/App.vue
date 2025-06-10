<template>
  <div class="app">
    <h1>Vote for Your Favorite</h1>
    <div class="buttons">
      <button @click="vote('Cats')">Cats</button>
      <button @click="vote('Dogs')">Dogs</button>
    </div>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: ''
    };
  },
  methods: {
    async vote(option) {
      try {
        const response = await fetch(`/vote?vote=${option}`, { method: 'POST' });

        if (!response.ok) {
          throw new Error(`Failed with status ${response.status}`);
        }

        const text = await response.text();
        this.message = text;
      } catch (err) {
        console.error(err);
        this.message = 'Failed to submit your vote. Please try again.';
      }
    }
  }
}
</script>

<style scoped>
.app {
  max-width: 600px;
  margin: 2rem auto;
  text-align: center;
  font-family: Arial, sans-serif;
}

.buttons {
  margin-top: 1rem;
}

button {
  margin: 0 10px;
  padding: 10px 20px;
  font-size: 1.2rem;
  cursor: pointer;
}
</style>
