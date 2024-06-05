import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";
import {AuthenticationUserStates, Recepie, Recepies} from '../../types/global'
import { updateRecepie } from "../../api/Yumms/RecepieUpdate";
import { toastNotification } from "../toast"

interface IndexProps {
    recepies: Recepie[];
}

const Index: React.FC<IndexProps> = ({ recepies }) => {
    const updateRecepie = async (userID: string, title: string, made: true, rating: number) => {
        const recepie: Recepie = {
            userID: userID,
            title: title,
            made: false,
            rating: 0
        };


        toastNotification('You have updated your recepie status', 'is-success');

    };

    return (
        <div>
            <div className="box cta">
                <p className="has-text-centered">
                    <span className="tag is-primary">Your Recipes</span>
                </p>
            </div>
            <section className="container is-variable is-4">
                <Splide options={{
                    perPage: 3,
                    arrows: false,
                }}>
                    {recepies.map((recepie) => (
                        <SplideSlide key={recepie.url}> {/* Added key for unique identification */}
                            <div className="card is-shady column is-10">
                                <div className="card-image">
                                    <figure className="image is-3by2">
                                        <img src={recepie.imageUrl} alt={recepie.title} className="modal-button"
                                             data-target="modal-image2"/>
                                    </figure>
                                </div>
                                <div className="card-content">
                                    <div className="content">
                                        <h5>{recepie.title}</h5>
                                        <div className="columns buttons is-mobile is-centered is-one-quarter">
                                            <a href={recepie.url} target="_blank" rel="noreferrer">
                                                <span className="button is-info modal-button column is-narrow"
                                                      data-target="modal-image2">Link</span>
                                            </a>
                                            <span className="button is-light column is-narrow"
                                                  data-target="modal-image2">Submit Review</span>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </SplideSlide>
                    ))}
                </Splide>
            </section>
        </div>
    );
}
export default Index;