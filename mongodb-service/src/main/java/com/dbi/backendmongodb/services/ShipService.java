package com.dbi.backendmongodb.services;

import com.dbi.backendmongodb.model.CrewMember;
import com.dbi.backendmongodb.model.Ship;
import com.dbi.backendmongodb.persistence.ShipRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@Transactional
public class ShipService {

    @Autowired
    private ShipRepository shipRepository;

    public void getShip(Ship ship) {}

    public void addShip(Ship ship) {}

    public void updateShip(Ship ship) {}

    public void removeShip(Ship ship) {}

    public void addCrewMemberToShip(CrewMember crewMember) {}
}
