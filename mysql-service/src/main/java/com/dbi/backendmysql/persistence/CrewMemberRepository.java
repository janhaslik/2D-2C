package com.dbi.backendmysql.persistence;

import com.dbi.backendmysql.model.CrewMember;
import org.springframework.data.repository.CrudRepository;

public interface CrewMemberRepository extends CrudRepository<CrewMember, Long> {
}
