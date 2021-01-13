<template>
  <form id="create_game">
    <h2>Create Game</h2>
    <div>opponent (hex address or leave it empty):</div>
    <input type="text" name="opponent" value="" />
    <input type="button" v-on:click="create_game" value="create" />
  </form>
</template>

<style scoped>
#create_game {
  margin: 50px auto;
}
#create_game input {
  font-size: 20px;
}
</style>

<script>
export default {
  data() {
    return {
      selected: "",
    };
  },
  methods: {
    async create_game() {
      const type = { type: "game" };
      const body = {};
      await this.$store.dispatch("entitySubmit", { ...type, body });
      await this.$store.dispatch("entityFetch", type);
    },
  },
};
</script>