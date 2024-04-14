package ru.bookinbl.search.myservice.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
@Entity
@Table(name = "offers")
public class Offer {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;

    @Column(name = "name")
    private String name;

    @Column(name = "cost")
    private Integer cost;

    @Column(name = "house")
    private Integer house;

    @Column(name = "rating")
    private Integer rating;

    @Column(name = "beds_count")
    private Integer bedsCount;

    @Column(name = "short_description")
    private Integer shortDescription;

    @ManyToOne
    @JoinColumn(name = "street_id")
    private Street street;

    @ManyToOne
    @JoinColumn(name = "type_of_housing_id")
    private Type type;
}
