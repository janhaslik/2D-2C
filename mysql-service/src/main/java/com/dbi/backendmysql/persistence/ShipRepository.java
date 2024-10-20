package com.dbi.backendmysql.persistence;

import com.dbi.backendmysql.model.Ship;
import org.springframework.data.repository.CrudRepository;

public interface ShipRepository extends CrudRepository<Ship, Integer> {
}
