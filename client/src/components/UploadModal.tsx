import React, { useState, ChangeEvent, FormEvent } from 'react';
import Modal from "./Modal";
import {Simulate} from "react-dom/test-utils";
import error = Simulate.error;

interface UploadModalProps {
    isOpen: boolean;
    onRequestClose: () => void;
}

const UploadModal: React.FC<UploadModalProps> = ({ isOpen, onRequestClose }) => {
    const [title, setTitle] = useState<string>('');
    const [body, setBody] = useState<string>('');
    const [image, setImage] = useState<File | null>(null);
    const [imageUrl, setImageUrl] = useState<string>('');

    const handleImageChange = (e: ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            setImage(e.target.files[0]);
        }
    };

    const handleImageUpload = async () => {
        if (!image) return;

        const formData = new FormData();
        formData.append('image', image);

        try {
            const response = await fetch('https://dummy.api/uploadImg', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: formData,
            });
            if (!response.ok){
                throw new Error('Network response was not ok')
            }
            const data = await response.json();

            setImageUrl(data.url);
        } catch (error) {
            console.error('Error uploading image:', error);
        }
    };

    const handleSubmit = (e: FormEvent) => {
        e.preventDefault();
        // Handle form submission logic here
        console.log('Title:', title);
        console.log('Body:', body);
        console.log('Image URL:', imageUrl);
        onRequestClose();
    };

    return (
        <Modal isOpen={isOpen} onClose={onRequestClose}>
            <h2>Upload Content</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Title:</label>
                    <input type="text" value={title} onChange={(e) => setTitle(e.target.value)} required />
                </div>
                <div>
                    <label>Body:</label>
                    <textarea value={body} onChange={(e) => setBody(e.target.value)} required />
                </div>
                <div>
                    <label>Image:</label>
                    <input type="file" accept="image/*" onChange={handleImageChange} />
                    <button type="button" onClick={handleImageUpload}>Upload Image</button>
                </div>
                {imageUrl && <img src={imageUrl} alt="Uploaded" style={{ maxWidth: '100%' }} />}
                <button type="submit">Submit</button>
            </form>
        </Modal>
    );
};

export default UploadModal;