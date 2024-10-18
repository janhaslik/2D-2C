package com.dbi.backendmongodb.model;

import com.fasterxml.jackson.annotation.JsonIgnore;
import lombok.*;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.sql.Date;
import java.util.List;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Document(collection = "ships")
public class Ship{

    @Id
    private int shipnr;
    private String name;
    private int owner;
    private String type;
    private String image;
    private String currentvalue;
    private Date year;

    private List<CrewMember> crewmembers;
}
