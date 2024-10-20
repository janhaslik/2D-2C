package com.dbi.backendmysql.model;

import com.fasterxml.jackson.annotation.JsonIgnore;
import jakarta.persistence.*;
import lombok.*;

@Getter
@Setter
@Builder
@NoArgsConstructor
@AllArgsConstructor
@Entity
public class CrewMember {

    @Id
    @GeneratedValue
    private int crewmemberid;
    private String name;
    private String role;

    @ManyToOne
    @JsonIgnore
    @JoinColumn(name = "shipnr")
    private Ship ship;
}
