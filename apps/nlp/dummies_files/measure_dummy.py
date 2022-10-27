import numpy as np
from numpy.linalg import norm

def cosine_sim(a1, a2):
    """
    Input - (numpy.ndarray) shape of (1, feature_dim) embedded vector of a sentence
    out - (numpy.float64) cosine sim of two vectors
    """
    return np.dot(a1, a2) / (norm(a1) * norm(a2))