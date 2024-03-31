package ru.bookinbl.search.myservice.storage;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import ru.bookinbl.search.myservice.model.Booking;

@Repository
public interface BookingRepository extends JpaRepository<Booking, Integer> {
    Page<Booking> findAllByCityAndRatingIsGreaterThanOrRating(String city, Integer rating1, Integer rating2, Pageable page);

    Page<Booking> findAllByRatingIsGreaterThanOrRating(Integer rating1, Integer rating2, Pageable page);
}
