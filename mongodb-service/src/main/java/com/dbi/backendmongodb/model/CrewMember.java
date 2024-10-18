package com.dbi.backendmongodb.model;

import lombok.*;
import org.springframework.data.annotation.Id;

@Getter
@Setter
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CrewMember {

    @Id
    private int crewmemberid;
    private String name;
    private String role;

}

