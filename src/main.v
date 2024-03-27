module main

const voting_age = 18

struct User {
	name string @[required]
	age  u8     @[required]
mut:
	is_registered bool
}

fn (mut user User) register() !bool {
	if user.age >= voting_age {
		user.is_registered = true
		return true
	} else {
		return error("user '${user.name}' is too young to register for voting")
	}
}

fn main() {
	mut john := User{
		name: 'John Doe'
		age: 21
	}
	println(john)

	_ := john.register() or {
		println('error: ${err}')
		false
	}
	println(john)
}
