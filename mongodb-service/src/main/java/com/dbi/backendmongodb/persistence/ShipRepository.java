package com.dbi.backendmongodb.persistence;

import com.dbi.backendmongodb.model.Ship;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ShipRepository extends MongoRepository<Ship, String> {
}
