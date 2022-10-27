from sklearn.metrics.pairwise import cosine_similarity

def cosine_sim(embed_vec1, embed_vec2):
    """
    Input - (numpy.ndarray) shape of (768, ) embedded vector of a sentence
    out - (numpy.float64) cosine sim of two vectors
    """
    return cosine_similarity([embed_vec1], [embed_vec2]) 