import React, { useEffect, useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";
import { Recepie } from '../../types/global';
import { toastNotification } from "../toast";
import { ReviewCard } from "./ReviewModal";

interface IndexProps {
    recepies: Recepie[];
}

const Index: React.FC<IndexProps> = ({ recepies }) => {
    const [showModal, setShowModal] = useState<boolean>(false);
    const [selectedRecepie, setSelectedRecepie] = useState<Recepie | null>(null);

    const handleReviewClick = (recepie: Recepie) => {
        setSelectedRecepie(recepie);
        setShowModal(true);
    };

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
            {showModal && selectedRecepie && (
                <ReviewCard open={showModal} onClose={() => setShowModal(false)} recepie={selectedRecepie} />
            )}
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
                        <SplideSlide key={recepie.url}>
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
                                            <span className="button is-success column is-narrow"
                                                  data-target="modal-image2" onClick={() => handleReviewClick(recepie)}>Submit Review</span>
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