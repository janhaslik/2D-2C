package main

type Seeder struct {
	seed int32
}

func NewSeeder() Seeder {
	return Seeder{
		seed: 31,
	}
}

func (seeder *Seeder) SeedCrewMembers() ([]CrewMember, error) {
	// cat := make(map[string]string)

	// cat["cat1"] = "The cat is wet"
	// cat["cat2"] = "The cat is fat"
	// cat["cat3"] = "The cat is meow meow"

	// return cat, nil
	return nil, nil
}

func (seeder *Seeder) SeedShips() ([]Ship, error) {
	// cat := make(map[string]string)

	// cat["cat1"] = "The cat is wet"
	// cat["cat2"] = "The cat is fat"
	// cat["cat3"] = "The cat is black"

	// return cat, nil
	return nil, nil
}
