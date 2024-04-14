package ru.bookinbl.search.myservice.storage;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import ru.bookinbl.search.myservice.model.Street;

@Repository
public interface StreetRepository extends JpaRepository<Street, Integer> {
}
