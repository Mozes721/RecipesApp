import React, { useEffect, useState } from "react";

interface ModalProps {
    isOpen: boolean;
    onClose: () => void;
}

const TableTypes: React.FC<ModalProps> = ({ isOpen, onClose }) => {
    return isOpen ? (
        <div className="content">
            <h2 className="has-text-grey">Meal Types</h2>
            <table>
                <thead>
                <th>1</th>
                <th>2</th>
                </thead>
                <tbody>
                <tr>
                    <th>main course</th>
                    <th>side dish</th>
                </tr>
                <tr>
                    <th>dessert</th>
                    <th>appetizer</th>
                </tr>
                <tr>
                    <th>salad</th>
                    <th>bread</th>
                </tr>
                <tr>
                    <th>breakfast</th>
                    <th>soup</th>
                </tr>
                <tr>
                    <th>beverage</th>
                    <th>sauce</th>
                </tr>
                <tr>
                    <th>marinade</th>
                    <th>fingerfood</th>
                </tr>
                <tr>
                    <th>snack</th>
                    <th>drink</th>
                </tr>
                </tbody>
            </table>
        </div>
    ) : null;
};

export default TableTypes;