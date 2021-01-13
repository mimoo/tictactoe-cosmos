<template>
  <div>
    <h2>GAMES</h2>

    <div class="game" v-for="game in games">
      <h3>game id: {{ game.id }}</h3>
      <p>
        created by: <strong>{{ game.creator }}</strong> for
        <span v-if="game.opponent == 'no'">anyone</span>
        <span v-else>{{ game.opponent }}</span>
      </p>
      <p>status: {{ game.status }}</p>
      <pre>
{{ game.state.substring(0, 3) }}<br>
{{ game.state.substring(3, 6) }}<br>
{{ game.state.substring(6, 9) }}</pre>
    </div>
  </div>
</template>

<style scoped>
.game {
  border: 1px solid black;
  padding: 20px;
}
span {
  font-weight: bold;
}
pre {
  padding: 5px;
  background-color: rgb(198, 201, 194);
  font-size: 20px;
}
</style>

<script>
export default {
  computed: {
    games() {
      let games = this.$store.state.data["game"];
      return games || [];
    },
  },
};
</script>
