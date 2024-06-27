import React, { useState } from "react";
import { AuthenticationUserStates, Recepie, ReviewedRecepie } from '../../types/global';
import { Star} from "phosphor-react";
import { updateRecepie } from "../../api/Yumms/RecepieUpdate";
import { deleteRecepie } from "../../api/Yumms/RecepieDelete";
import { useSelector } from "react-redux";

interface ModalProps {
    open: boolean;
    onClose: () => void;
    recepie: Recepie;
}

export const ReviewCard: React.FC<ModalProps> = ({ open, onClose, recepie }) => {
    const userID = useSelector((state: AuthenticationUserStates) => state.userID);
    const [rating, setRating] = useState<number>(recepie.rating ?? 0);

    if (!open) return null;

    const handleSaveChanges = async (id: string | undefined) => {
        const reviewed: ReviewedRecepie = {
            id: id,
            rating: rating
        };

        const response = await updateRecepie(reviewed, userID);
        if (response) {
            onClose();
            window.location.href = '/yumms';
        }

    };

    const handleDeleteChanges = async () => {
        const response = await deleteRecepie(userID, recepie.id);
        if (response) {
            onClose();
            window.location.href = '/yumms';
        }
    };

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
                        {[1, 2, 3, 4, 5].map((starIndex) => (
                            <button
                                key={starIndex}
                                className="button p-2 is-ghost p-0"
                                onClick={() => setRating(starIndex)}
                            >
                                <Star size={40} weight={rating >= starIndex ? "fill" : "regular"} />
                            </button>
                        ))}
                    </div>
                </section>
                <footer className="modal-card-foot">
                    <div className="buttons">
                        <button className="button is-success"  onClick={() => handleSaveChanges(recepie.id)}>Save changes</button>
                        <button className="button is-danger" onClick={handleDeleteChanges}>Remove recepie</button>
                        <button className="button" onClick={onClose}>Cancel</button>
                    </div>
                </footer>
            </div>
        </div>
    );
};
