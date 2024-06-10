import React from "react";
import { Recepie } from '../../types/global';

interface ModalProps {
    open: boolean;
    onClose: () => void;
    recepie: Recepie;
}

export const ReviewCard: React.FC<ModalProps> = ({ open, onClose, recepie }) => {
    if (!open) return null;

    return (
        <div className={`modal ${open ? "is-active" : ""}`}>
            <div className="modal-background" onClick={onClose}></div>
            <div className="modal-card">
                <header className="modal-card-head">
                    <p className="modal-card-title">{recepie.title}</p>
                    <button className="delete" aria-label="close" onClick={onClose}></button>
                </header>
                <section className="modal-card-body">
                    <div className="mb-10">
                        <button className="button is-ghost p-0">
                            <img className="image is-fullwidth" src="yofte-assets/elements/star-gold.svg" alt=""/>
                        </button>
                        <button className="button is-ghost p-0">
                            <img className="image is-fullwidth" src="yofte-assets/elements/star-gold.svg" alt=""/>
                        </button>
                        <button className="button is-ghost p-0">
                            <img className="image is-fullwidth" src="yofte-assets/elements/star-gold.svg" alt=""/>
                        </button>
                        <button className="button is-ghost p-0">
                            <img className="image is-fullwidth" src="yofte-assets/elements/star-gold.svg" alt=""/>
                        </button>
                        <button className="button is-ghost p-0">
                            <img className="image is-fullwidth" src="yofte-assets/elements/star-gray.svg" alt=""/>
                        </button>
                    </div>
                </section>
                <footer className="modal-card-foot">
                    <div className="buttons">
                        <button className="button is-success">Save changes</button>
                        <button className="button" onClick={onClose}>Cancel</button>
                    </div>
                </footer>
            </div>
        </div>
    );
};
