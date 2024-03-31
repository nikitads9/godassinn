package ru.bookinbl.search.myservice.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.http.client.HttpComponentsClientHttpRequestFactory;
import org.springframework.stereotype.Service;
import org.springframework.web.util.DefaultUriBuilderFactory;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Map;
import java.util.StringJoiner;

@Service
public class OtherServiceClient extends BaseClient {
    private static final String START = "start";
    private static final String END = "end";
    public static final String DATETIME_FORMAT = "yyyy-MM-dd HH:mm:ss";
    public static final DateTimeFormatter FORMATTER = DateTimeFormatter.ofPattern(DATETIME_FORMAT);

    @Autowired
    public OtherServiceClient(@Value("http://bookings:3000") String serverUrl, RestTemplateBuilder builder) {
        super(
                builder
                        .uriTemplateHandler(new DefaultUriBuilderFactory(serverUrl))
                        .requestFactory(HttpComponentsClientHttpRequestFactory::new)
                        .build()
        );
    }

    public ResponseEntity<Object> getBookingsWithTime(LocalDateTime rangeStart, LocalDateTime rangeEnd) {

        String stringStart = rangeStart.format(FORMATTER);
        String stringEnd = rangeEnd.format(FORMATTER);

        Map<String, Object> parameters = Map.of(
                START, stringStart,
                END, stringEnd
        );

        StringJoiner pathBuilder = new StringJoiner("&", "/bookings?start={start}&end={end}\"", "");

        String path = pathBuilder.toString();
        return makeAndSendRequest(HttpMethod.GET, path, null, parameters, null);
    }
}