package ru.bookinbl.search.myservice.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.stereotype.Service;
import ru.bookinbl.search.myservice.model.Offer;
import ru.bookinbl.search.myservice.storage.OfferRepository;

import java.util.List;

@Service
@RequiredArgsConstructor
public class OfferServiceImpl implements OfferService {

    private final OfferRepository offerRepository;

    @Override
    public List<Offer> getBookings(String city, int rating, int from, int size) {

        PageRequest page = PageRequest.of(from > 0 ? from / size : 0, size);
        Page<Offer> events;

        if (!city.isBlank()) {
            events = offerRepository.findAllByCityAndRatingIsGreaterThanOrRating(city, rating, rating, page);
        } else {
            events = offerRepository.findAllByRatingIsGreaterThanOrRating(rating, rating, page);
        }
        return events.getContent();
    }
}
