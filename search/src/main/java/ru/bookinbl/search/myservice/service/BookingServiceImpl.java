package ru.bookinbl.search.myservice.service;

import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.stereotype.Service;
import ru.bookinbl.search.myservice.model.Booking;
import ru.bookinbl.search.myservice.storage.BookingRepository;

import java.util.List;

@Service
@RequiredArgsConstructor
public class BookingServiceImpl implements BookingService {

    private final BookingRepository bookingRepository;

    @Override
    public List<Booking> getBookings(String city, int rating, int from, int size) {

        PageRequest page = PageRequest.of(from > 0 ? from / size : 0, size);
        Page<Booking> events;

        if (!city.isBlank()) {
            events = bookingRepository.findAllByCityAndRatingIsGreaterThanOrRating(city, rating, rating, page);
        } else {
            events = bookingRepository.findAllByRatingIsGreaterThanOrRating(rating, rating, page);
        }
        return events.getContent();
    }
}
