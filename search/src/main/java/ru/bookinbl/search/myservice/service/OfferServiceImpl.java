package ru.bookinbl.search.myservice.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.stereotype.Service;
import ru.bookinbl.search.myservice.error.exception.EntityNotFoundException;
import ru.bookinbl.search.myservice.model.Offer;
import ru.bookinbl.search.myservice.model.Street;
import ru.bookinbl.search.myservice.storage.OfferRepository;
import ru.bookinbl.search.myservice.storage.StreetRepository;

import java.util.List;

@Service
@RequiredArgsConstructor
public class OfferServiceImpl implements OfferService {

    private final OfferRepository offerRepository;
    private final StreetRepository streetRepository;

    @Override
    public List<Offer> getBookings(Integer streetId, int rating, int from, int size) {

        PageRequest page = PageRequest.of(from > 0 ? from / size : 0, size);
        Page<Offer> events;

        if (streetId!=null) {
            Street street = streetRepository.findById(streetId).orElseThrow(() -> new EntityNotFoundException("Введённой улицы не существует"));
            events = offerRepository.findAllByStreetAndRatingIsGreaterThanOrRating(street, rating, rating, page);
        } else {
            events = offerRepository.findAllByRatingIsGreaterThanOrRating(rating, rating, page);
        }
        return events.getContent();
    }
}
