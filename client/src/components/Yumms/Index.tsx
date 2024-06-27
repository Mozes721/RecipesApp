import React, { useState } from "react";
import {Splide, SplideSlide} from "@splidejs/react-splide";
import "@splidejs/splide/dist/css/splide.min.css";
import { Recepie } from '../../types/global';
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
                                        {recepie.made ?
                                            ( <h6 className="subtitle is-6">
                                                You have already submitted review on this recepie
                                            </h6> )
                                            :
                                            null
                                        }
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