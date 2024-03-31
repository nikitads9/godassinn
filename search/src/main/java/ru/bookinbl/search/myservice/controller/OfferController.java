package ru.bookinbl.search.myservice.controller;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.format.annotation.DateTimeFormat;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import ru.bookinbl.search.myservice.model.Offer;
import ru.bookinbl.search.myservice.service.OfferService;
import ru.bookinbl.search.myservice.service.OtherServiceClient;

import javax.validation.constraints.Positive;
import javax.validation.constraints.PositiveOrZero;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.List;

@RestController
@RequiredArgsConstructor
@Slf4j
@Validated
public class OfferController {

    private final OfferService offerService;
    private final OtherServiceClient otherServiceClient;

    public static final String DATETIME_FORMAT = "yyyy-MM-dd HH:mm:ss";
    public static final DateTimeFormatter FORMATTER = DateTimeFormatter.ofPattern(DATETIME_FORMAT);

    @GetMapping("/MyBookings")
    public List<Offer> getBookings(@RequestParam(required = false, defaultValue = "") String city,
                                   @RequestParam(required = false, defaultValue = "0") Integer rating,
                                   @RequestParam(defaultValue = "0") @PositiveOrZero int from,
                                   @RequestParam(defaultValue = "10") @Positive int size) {
        log.debug("Вызван метод getBookings");
        return offerService.getBookings(city, rating, from, size);
    }

    @GetMapping("/MyBookings/time")
    public ResponseEntity<Object> getBookingsWithTime(@RequestParam(required = false, defaultValue = "#{T(java.time.LocalDateTime).now()}") @DateTimeFormat(pattern = DATETIME_FORMAT) LocalDateTime rangeStart,
                                                      @RequestParam(required = false) @DateTimeFormat(pattern = DATETIME_FORMAT) LocalDateTime rangeEnd) {
        log.debug("Вызван метод getBookingsWithTime");
        if (rangeEnd == null) {
            rangeEnd = rangeStart.plusDays(14);
        }
        return otherServiceClient.getBookingsWithTime(rangeStart, rangeEnd);
    }

}
