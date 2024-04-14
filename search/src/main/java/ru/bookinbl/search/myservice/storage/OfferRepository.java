package ru.bookinbl.search.myservice.storage;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import ru.bookinbl.search.myservice.model.Offer;
import ru.bookinbl.search.myservice.model.Street;

@Repository
public interface OfferRepository extends JpaRepository<Offer, Integer> {
    Page<Offer> findAllByStreetAndRatingIsGreaterThanOrRating(Street street, Integer rating, Integer rating2, Pageable pageable);

    Page<Offer> findAllByRatingIsGreaterThanOrRating(Integer rating1, Integer rating2, Pageable page);
}
