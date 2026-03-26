package main

func spawnEntity(Entity Entity) {
	game.Entities = append(game.Entities, Entity)
	Entity.Start()
}
