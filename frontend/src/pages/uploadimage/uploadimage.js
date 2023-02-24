import React, { useState } from 'react';
import axios from 'axios';

function ImageUploader() {
  const [image, setImage] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleChange = (event) => {
    setImage(event.target.files[0]);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);
    const formData = new FormData();
    formData.append("image", image);
    try {
      await axios.post('http://localhost:8080/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      alert('Image uploaded successfully!');
    } catch (err) {
      console.error(err);
      alert('Error uploading image.');
    }
    setLoading(false);
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="file" onChange={handleChange} />
      <button type="submit" disabled={loading}>
        Upload
      </button>
    </form>
  );
}

export default ImageUploader;
