module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "start", fields: ["gameID", ] },
		{ type: "move", fields: ["gameID", "position", ] },
		{ type: "game", fields: ["status", "opponent", "move_num", "state", ] },
  ],
};
