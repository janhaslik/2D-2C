package com.dbi.backendmysql.model;

import com.fasterxml.jackson.annotation.JsonIgnore;
import jakarta.persistence.*;
import lombok.*;

import java.sql.Date;
import java.util.List;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Entity
public class Ship{

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int shipnr;
    private String name;
    private int owner;
    private String type;
    private String image;
    private String currentvalue;
    private Date year;

    @OneToMany(cascade = CascadeType.PERSIST, mappedBy = "ship")
    private List<CrewMember> crewmembers;
}
