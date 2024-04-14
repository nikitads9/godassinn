package ru.bookinbl.search.myservice.error.handler;

import org.springframework.http.HttpStatus;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.web.HttpRequestMethodNotSupportedException;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.method.annotation.MethodArgumentTypeMismatchException;
import org.springframework.web.server.MethodNotAllowedException;
import org.springframework.web.servlet.NoHandlerFoundException;
import ru.bookinbl.search.myservice.error.exception.EntityNotFoundException;
import ru.bookinbl.search.myservice.error.model.ApiError;

@RestControllerAdvice
public class MyHandler {

    @ExceptionHandler(value = {
            HttpMessageNotReadableException.class,
            NoHandlerFoundException.class,
            MethodNotAllowedException.class,
            HttpRequestMethodNotSupportedException.class,
            NumberFormatException.class,
            MethodArgumentNotValidException.class,
            MethodArgumentTypeMismatchException.class,
    })
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    public ApiError handleBadRequest(final Exception e) {
        return new ApiError(null, e.getMessage(), "Incorrectly made request", HttpStatus.BAD_REQUEST);
    }
    @ExceptionHandler(value = {
            EntityNotFoundException.class
    })
    @ResponseStatus(HttpStatus.NOT_FOUND)
    public ApiError handleNotfoundException(final Exception exception) {
        return new ApiError(null, exception.getMessage(), "The required object was not found", HttpStatus.NOT_FOUND);
    }
}
