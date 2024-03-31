package ru.bookinbl.search.myservice.service;

import ru.bookinbl.search.myservice.model.Offer;

import java.util.List;

public interface OfferService {

    List<Offer> getBookings(String city, int rating, int from, int size);
}
