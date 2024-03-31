package ru.bookinbl.search.myservice.service;

import ru.bookinbl.search.myservice.model.Booking;

import java.util.List;

public interface BookingService {

    List<Booking> getBookings(String city, int rating, int from, int size);
}
